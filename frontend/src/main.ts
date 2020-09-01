import App from './App.svelte';
import runtime from '@wailsapp/runtime';

declare global {
  interface MergeElement {
    name: string;
    data: string;
    pages: string;
  }

  interface Window {
    backend: {
      MergePDFs: (arg0: {
        mergeElements: Array<MergeElement>;
      }) => Promise<string>;
    };
  }
}

runtime.Init(() => {
  new App({
    target: document.body,
    props: {}
  });
});
