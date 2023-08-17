const btns_close = document.querySelectorAll('.btn-close');
const btn_symbols = document.querySelector('#btn-symbol');
const btn_errors = document.querySelector('#btn-error');
const btn_reports = document.querySelector('#btn-report');
const btn_run = document.querySelector('#btn-run');
const btn_clear = document.querySelector('#btn-clear');
const terminal = document.querySelector('#terminal');
const editor = document.querySelector('#editor');

btns_close.forEach(btn => {
    btn.addEventListener('click', () => {
        $(btn).parent().parent().fadeOut(300);
    });
});

btn_symbols.addEventListener('click', (e) => {
    e.preventDefault();
    $('#symbol-modal').fadeIn(300);
});

btn_errors.addEventListener('click', (e) => {
    e.preventDefault();
    $('#error-modal').fadeIn(300);
});

btn_reports.addEventListener('click', (e) => {
    e.preventDefault();
    $('#report-modal').fadeIn(300);
});

/* Enviar datos a API golang */
btn_run.addEventListener('click', async (e) => {
    e.preventDefault();

    let data = {
        "code": editor.value
    }
    
    let response = await fetch('http://localhost:3000/run', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
    });

    if (!response.ok) {
        console.log('Error al enviar los datos');
        return;
    }

    let message = await response.json();
    terminal.value +=  message.message;
});

/* Limpiar termianl */
btn_clear.addEventListener('click', (e) => {
    e.preventDefault();
    terminal.value = '';
})
