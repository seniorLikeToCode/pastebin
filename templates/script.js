const pb = document.getElementById('paste-here');
// path here innerHTML of pb 

function pasteFromClipboard() {
    navigator.clipboard.readText().then(function (text) {
        // console.log('Pasted text: ', text);
        pb.textContent = text;
        const result = hljs.highlightAuto(text);
        pb.className = result.language; // set language class for highlight.js
        hljs.highlightElement(pb);

    }).catch(function (err) {
        console.error('Could not read text from clipboard: ', err);
    });
}

function copyToClipboard(text) {
    navigator.clipboard.writeText(text).then(function () {
        console.log('Text copied to clipboard');
    }).catch(function (err) {
        console.error('Could not copy text: ', err);
    });
}

// Usage example
document.addEventListener('keydown', function (event) {
    console.log('Key pressed: ', event.key, ' Ctrl key pressed: ', event.ctrlKey)
    if (event.ctrlKey && (event.key === 'v' || event.key === 'V')) {
        event.preventDefault();
        document.querySelector('code').removeAttribute('data-highlighted');
        pasteFromClipboard();
    }
});