<script>
  import Controls from './Controls.svelte';
  import Preview from './Preview.svelte';

  let pdfFiles;
  let mergedPDF = '';

  async function merge() {
    let mergeElements = [];

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
      mergedPDF = null;
    }
  }
  function readFileAsync(file) {
    return new Promise((resolve, reject) => {
      let reader = new FileReader();

      reader.onload = () => {
        resolve(reader.result);
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
    position: absolute;
    left: 0px;
    width: 50%;
    height: 100%;
    background: #3c476a; /* Old browsers */
    background: -moz-linear-gradient(
      top,
      #3c476a 0%,
      #252e48 100%
    ); /* FF3.6-15 */
    background: -webkit-linear-gradient(
      top,
      #3c476a 0%,
      #252e48 100%
    ); /* Chrome10-25,Safari5.1-6 */
    background: linear-gradient(
      to bottom,
      #3c476a 0%,
      #252e48 100%
    ); /* W3C, IE10+, FF16+, Chrome26+, Opera12+, Safari7+ */
    filter: progid:DXImageTransform.Microsoft.gradient( startColorstr='#3c476a', endColorstr='#252e48',GradientType=0 ); /* IE6-9 */
  }

  .preview {
    text-align: center;
    position: absolute;
    right: 0px;
    width: 50%;
    overflow: auto;
    background: #252e48; /* Old browsers */
    background: -moz-linear-gradient(
      top,
      #252e48 0%,
      #3c476a 100%
    ); /* FF3.6-15 */
    background: -webkit-linear-gradient(
      top,
      #252e48 0%,
      #3c476a 100%
    ); /* Chrome10-25,Safari5.1-6 */
    background: linear-gradient(
      to bottom,
      #252e48 0%,
      #3c476a 100%
    ); /* W3C, IE10+, FF16+, Chrome26+, Opera12+, Safari7+ */
    filter: progid:DXImageTransform.Microsoft.gradient( startColorstr='#252e48', endColorstr='#3c476a',GradientType=0 ); /* IE6-9 */
  }

  #mergeBtn {
    z-index: 9999;
  }
</style>

<main>
  <section class="container">
    <div class="controls">
      <Controls bind:pdfFiles />
    </div>

    <button
      id="mergeBtn"
      on:click={async () => {
        await merge();
      }}>
      merge
    </button>

    <div class="preview">
      <Preview pdfData="${mergedPDF}" />
    </div>
  </section>
</main>
