<script lang="ts">
  import Controls from './Controls.svelte';
  import Preview from './Preview.svelte';
</script>

<style>
  div {
    height: 100%;
  }

  .container {
    width: 100%;
    display: flex;
    align-items: stretch;
  }

  .controls {
    background-color: #ff9e2c;
    position: absolute;
    left: 0px;
    width: 50%;
  }

  .preview {
    background-color: #b6701e;
    position: absolute;
    right: 0px;
    width: 50%;
  }
</style>

<main>
  <section class="container">
    <div class="controls">
      <Controls />
    </div>
    <div class="preview">
      <Preview />
    </div>
  </section>
</main>
