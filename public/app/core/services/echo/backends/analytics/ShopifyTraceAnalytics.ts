import { EchoBackend, EchoEventType, getDataSourceSrv, PageviewEchoEvent } from '@grafana/runtime';

export class ShopifyAnalyticsBackend implements EchoBackend<PageviewEchoEvent, {}> {
  supportedEvents = [EchoEventType.Pageview];

  queuedEvents: PageviewEchoEvent[] = [];
  datasource: any = null;

  constructor(public options: {}) {
    this.loadDatasource();
  }

  loadDatasource = async () => {
    try {
      const ds = await getDataSourceSrv().get('Trace Search');
      this.datasource = ds;
      this.flushQueue();
    } catch (e) {
      setTimeout(this.loadDatasource, 1000);
    }
  };

  flushQueue = () => {
    this.queuedEvents.forEach(this.addEvent);
  };

  addEvent = (e: PageviewEchoEvent) => {
    if (!this.datasource) {
      this.queuedEvents.push(e);
      return;
    }

    const { userId } = e.meta;

    this.datasource.trackPageView(userId, `${window.location.origin}${e.payload.page}`);
  };

  // Not using Echo buffering, addEvent above sends events to GA as soon as they appear
  flush = () => {};
}
