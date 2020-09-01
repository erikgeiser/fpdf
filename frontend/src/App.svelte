<script lang="ts">
  import Controls from './Controls.svelte';
  import Preview from './Preview.svelte';

  let pdfFiles: File[];
  let mergedPDF: string = '';

  async function merge() {
    let mergeElements: Array<MergeElement> = [];

    for (let i = 0; i < pdfFiles.length; i++) {
      let file = pdfFiles[i];

      mergeElements.push({
        name: file.name,
        data: btoa(
          String.fromCharCode.apply(
            null,
            new Uint8Array(await readFileAsync(file))
          )
        ),
        pages: '1'
      });
    }
    try {
      mergedPDF = atob(await window.backend.MergePDFs({ mergeElements }));
    } catch (e) {
      alert(e);
    }
  }
  function readFileAsync(file: File): Promise<ArrayBuffer> {
    return new Promise((resolve, reject) => {
      let reader = new FileReader();

      reader.onload = () => {
        resolve(reader.result as ArrayBuffer);
      };

      reader.onerror = reject;

      reader.readAsArrayBuffer(file);
    });
  }
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
      <Controls bind:pdfFiles />
      <button
        on:click={async () => {
          await merge();
        }}>
        merge
      </button>
    </div>
    <div class="preview">
      <Preview pdfData="${mergedPDF}" />
    </div>
  </section>
</main>
