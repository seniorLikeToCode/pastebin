const pb = document.getElementById('paste-here');
const pbTextArea = document.getElementById('paste-here-textarea');
pb.contentEditable = true;

const url = window.location.href;
const parts = url.split('/');
const id = parts[parts.length - 1];

function isValidID(uid) {
    if (uid.length !== 6) return true;
    pb.contentEditable = false;
    return false;
}

function handleRedirect(uid) {
    const baseUrl = window.location.href; // The current URL or any base URL you have
    const newUrl = `${baseUrl}/${uid}`; // Appending the ID to the URL
    window.location.href = newUrl;  // Redirect to the new URL
}

async function getLinkContent(uid) {
    if (isValidID(uid)) return;
    const response = await fetch(`http://localhost:5000/api/v1/${uid}`, {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json',
        }
    });
    const data = await response.json();

    pbTextArea.classList.add('hidden');
    pb.classList.remove('hidden');

    pb.textContent = data.content;
    const result = hljs.highlightAuto(data.content);
    pb.className = result.language; // set language class for highlight.js
    hljs.highlightElement(pb);

}

getLinkContent(id);


function pasteFromClipboard() {
    navigator.clipboard.readText().then(function (text) {
        pbTextArea.value = text;

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
            content: pbTextArea.value,
        })
    });

    const data = await response.json();
    handleRedirect(data.id);
}


// Usage example
document.addEventListener('keydown', function (event) {
    // console.log('Key pressed: ', event.key, ' Ctrl key pressed: ', event.ctrlKey)
    if (event.ctrlKey && (event.key === 'v' || event.key === 'V')) {
        event.preventDefault();
        pb.removeAttribute('data-highlighted');
        pasteFromClipboard();
    }

    if (event.ctrlKey && (event.key === 's' || event.key === 'S')) {
        event.preventDefault();
        createLink();
    }
});