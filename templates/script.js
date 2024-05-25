const pb = document.getElementById('paste-here');

const url = window.location.href;
const parts = url.split('/');
const id = parts[parts.length - 1];


function isValidID(id) {
    return (id.length !== 6);
}

async function getLinkContent(id) {
    if (isValidID(id)) return;
    const response = await fetch(`http://localhost:5000/api/v1/${id}`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        }
    });
    const data = await response.json();
    pb.textContent = data.content;
    const result = hljs.highlightAuto(data.content);
    pb.className = result.language; // set language class for highlight.js
    hljs.highlightElement(pb);
}

getLinkContent(id);


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

async function createLink() {
    const response = await fetch('http://localhost:5000/api/v1/', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            content: pb.textContent,
        })
    });

    const data = await response.json();
    const baseUrl = window.location.href; // The current URL or any base URL you have
    const newUrl = `${baseUrl}/${data.id}`; // Appending the ID to the URL

    // Redirect to the new URL
    window.location.href = newUrl;
}




// Usage example
document.addEventListener('keydown', function (event) {
    // console.log('Key pressed: ', event.key, ' Ctrl key pressed: ', event.ctrlKey)
    if (event.ctrlKey && (event.key === 'v' || event.key === 'V')) {
        event.preventDefault();
        document.querySelector('code').removeAttribute('data-highlighted');
        pasteFromClipboard();
    }

    if (event.ctrlKey && (event.key === 's' || event.key === 'S')) {
        event.preventDefault();
        createLink();
    }
});