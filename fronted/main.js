const btns_close = document.querySelectorAll('.btn-close');
const btn_symbols = document.querySelector('#btn-symbol');
const btn_errors = document.querySelector('#btn-error');
const btn_reports = document.querySelector('#btn-report');

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