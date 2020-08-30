import App from './App.svelte';
import runtime from '@wailsapp/runtime';

declare global {
  interface Window {
    backend: any;
  }
}

runtime.Init(() => {
  new App({
    target: document.body,
    props: {}
  });
});
