<script>
  import pdflib from 'pdfjs-dist/es5/build/pdf';
  import pdfjsWorker from 'pdfjs-dist/es5/build/pdf.worker.entry';

  export let pdfData;

  async function render(data) {
    let canvasContainer = document.getElementById('pdf');
    if (!canvasContainer) {
      return;
    }

    canvasContainer.textContent = '';

    if (pdfData == '' || pdfData == null) {
      return;
    }

    var options = { scale: 1 };

    function renderPage(page) {
      var viewport = page.getViewport(options);
      var canvas = document.createElement('canvas');
      var ctx = canvas.getContext('2d');
      var renderContext = {
        canvasContext: ctx,
        viewport: viewport
      };

      canvas.height = viewport.height;
      canvas.width = viewport.width;

      canvasContainer.appendChild(canvas);

      page.render(renderContext);
    }

    function renderPages(pdfDoc) {
      for (var num = 1; num <= pdfDoc.numPages; num++)
        pdfDoc.getPage(num).then(renderPage);
    }

    let doc = await pdflib.getDocument({ data }).promise;
    renderPages(doc);
  }

  $: pdfData ? render(pdfData) : 'nothing';
</script>

<style>
</style>

<main>
  {#if pdfData == null}error{/if}
  <div id="pdf" />
</main>
