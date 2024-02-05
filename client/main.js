class ButtonGet extends HTMLButtonElement {
    constructor() {
        super();
        // Load wasm_exec.js
        this.wasmExecPromise = fetch('wasm_exec.js')
            .then(response => response.text())
            .then(script => {
                eval(script);
                return new Go();
            })
            .catch(error => {
                console.error('Error:', error);
            });

        // Load wasm binary
        this.wasmModulePromise = this.wasmExecPromise.then(go => {
            return fetch('http.wasm')
                .then(response => response.arrayBuffer())
                .then(bytes => WebAssembly.instantiate(bytes, go.importObject))
                .then(results => {
                    go.run(results.instance);
                    return http();
                })
                .catch(error => {
                    console.error('Error:', error);
                });
        });
    }

    connectedCallback() {
        this.addEventListener('click', () => {
            this.wasmModulePromise.then(client => {
                // Get the URL to fetch from the 'url' attribute
                let url = this.getAttribute('url');
                // Call a function from the WASM module to make the HTTP request
                // This assumes that your WASM module has a function named 'Get'
                let responsePromise = client.Get(url);
                responsePromise.then(response => {
                    console.log('Response:', response);
                        // Get the selector of the element to replace from the 'replace' attribute
                        let selector = this.getAttribute('replace');
                        // Find the element to replace
                        let oldElement = document.querySelector(selector);
                        if (oldElement === null) {
                            console.error('Element not found:', selector);
                            return;
                        }
                        
                        // Create a new div element
                        let newDiv = document.createElement('div');
                        // Set the innerHTML of the new div to the returned HTML
                        newDiv.innerHTML =  response.body;
                        // Replace the old element with the new div
                        oldElement.replaceWith(newDiv);
                });
            });
        });
    }
}

customElements.define('button-get', ButtonGet, { extends: 'button' });