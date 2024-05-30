const baseURL = 'http://13.51.139.149:5000/api/v1/';    //13.51.139.149
const pb = document.getElementById('paste-here');
const pbTextArea = document.getElementById('paste-here-textarea');
pb.contentEditable = true;

const url = window.location.href;
const parts = url.split('/');
const id = parts[parts.length - 1];

function isValidID(uid) {
    if (uid.length !== 6) return true;
    pb.contentEditable = false;
    pbTextArea.classList.add('hidden');
    pb.classList.remove('hidden');
    return false;
}

function handleRedirect(uid) {
    const baseUrl = window.location.href; // The current URL or any base URL you have
    const newUrl = `${baseUrl}/${uid}`; // Appending the ID to the URL
    window.location.href = newUrl;  // Redirect to the new URL
}

async function getLinkContent(uid) {
    if (isValidID(uid)) return;
    const response = await fetch(baseURL + uid, {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json',
        }
    });
    const data = await response.json();

    pb.textContent = data.content;
    const result = hljs.highlightAuto(data.content);
    pb.className = result.language; // set language class for highlight.js
    hljs.highlightElement(pb);
    hljs.initLineNumbersOnLoad();

}

getLinkContent(id);

async function readFromClipboard() {
    try {
        const text = await navigator.clipboard.readText();
        console.log('Clipboard text:', text);
        pbTextArea.value = text;
    } catch (err) {
        console.error('Failed to read clipboard contents:', err);
    }
}


async function reqPermissionForRead() {
    const result = await navigator.permissions.query({ name: 'clipboard-write' });
    if (result.state === 'granted' || result.state === 'prompt') {
        readFromClipboard();
    } else {
        console.error('Clipboard permission denied');
    }
}

function copyToClipboard(text) {
    navigator.clipboard.writeText(text).then(function () {
        console.log('Text copied to clipboard');
    }).catch(function (err) {
        console.error('Could not copy text: ', err);
    });
}

async function createLink() {
    const response = await fetch(baseURL, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            content: pbTextArea.value,
        })
    });

    const data = await response.json();
    handleRedirect(data.id);
}


// Usage example
document.addEventListener('keydown', function (event) {
    if (!isValidID(id)) return;
    console.log('Key pressed: ', event.key, ' Ctrl key pressed: ', event.ctrlKey)
    if (event.ctrlKey && (event.key === 'v' || event.key === 'V')) {
        event.preventDefault();
        pb.removeAttribute('data-highlighted');
        reqPermissionForRead();
    }

    if (event.ctrlKey && (event.key === 's' || event.key === 'S')) {
        event.preventDefault();
        createLink();
    }
});