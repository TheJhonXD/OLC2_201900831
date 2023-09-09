const btns_close = document.querySelectorAll('.btn-close');
const btn_symbols = document.querySelector('#btn-symbol');
const btn_errors = document.querySelector('#btn-error');
const btn_reports = document.querySelector('#btn-report');
const btn_run = document.querySelector('#btn-run');
const btn_clear = document.querySelector('#btn-clear');
const terminal = document.querySelector('#terminal');
const editor = document.querySelector('#editor');
const save_as = document.querySelector('#save-as');


/* *************************************************************************** */
let SAMPLE_PARSER =
    "parser grammar ExprParser;\n" +
    "options { tokenVocab=ExprLexer; }\n" +
    "\n" +
    "program\n" +
    "    : stat EOF\n" +
    "    | def EOF\n" +
    "    ;\n" +
    "\n" +
    //"foo : 'a' 'abc' 'a\\'b' '\\u34ab' 'ab\\ncd' ;\n" +
    "stat: ID '=' expr ';'\n" +
    "    | expr ';'\n" +
    "    ;\n" +
    "\n" +
    "def : ID '(' ID (',' ID)* ')' '{' stat* '}' ;\n" +
    "\n" +
    "expr: ID\n" +
    "    | INT\n" +
    "    | func\n" +
    "    | 'not' expr\n" +
    "    | expr 'and' expr\n" +
    "    | expr 'or' expr\n" +
    "    ;\n" +
    "\n" +
    "func : ID '(' expr (',' expr)* ')' ;"

let SAMPLE_LEXER =
    "// DELETE THIS CONTENT IF YOU PUT COMBINED GRAMMAR IN Parser TAB\n" +
    "lexer grammar ExprLexer;\n" +
    "\n" +
    "AND : 'and' ;\n" +
    "OR : 'or' ;\n" +
    "NOT : 'not' ;\n" +
    "EQ : '=' ;\n" +
    "COMMA : ',' ;\n" +
    "SEMI : ';' ;\n" +
    "LPAREN : '(' ;\n" +
    "RPAREN : ')' ;\n" +
    "LCURLY : '{' ;\n" +
    "RCURLY : '}' ;\n" +
    "\n" +
    "INT : [0-9]+ ;\n" +
    "ID: [a-zA-Z_][a-zA-Z_0-9]* ;\n" +
    "WS: [ \\t\\n\\r\\f]+ -> skip ;";

let SAMPLE_INPUT =
    "f(x,y) {\n" +
    "    a = 3+foo;\n" +
    "    x and y;\n" +
    "}";
/* *************************************************************************** */

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
    /* let data = {
        "grammar": "parser grammar ExprParser;\noptions { tokenVocab=ExprLexer; }\n\nprogram : stat EOF | def EOF;\nstat: ID '=' expr ';' | expr ';';\ndef : ID '(' ID (',' ID)* ')' '{' stat* '}' ;\nexpr: ID | INT | func | 'not' expr | expr 'and' expr | expr 'or' expr;\nfunc : ID '(' expr (',' expr)* ')' ;",
        "lexgrammar": "lexer grammar ExprLexer;\nAND : 'and' ;\nOR : 'or' ;\nNOT : 'not' ;\nEQ : '=' ;\nCOMMA : ',' ;\nSEMI : ';' ;\nLPAREN : '(' ;\nRPAREN : ')' ;\nLCURLY : '{' ;\nRCURLY : '}' ;\nINT : [0-9]+ ;\nID: [a-zA-Z_][a-zA-Z_0-9]* ;\nWS: [ \\t\\n\\r\\f]+ -> skip ;",
        "input": "f(x,y) {\n  a = 3;\n  x and y;\n}",
        "start": "program"
      } */
    let data = {
        "code": editor.value
    }
    prueba(data);
});

let prueba = async (data) =>{
    let response = await fetch('http://localhost:3000/tree', {
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
    if (message.result.svgtree != null) {
        let b64_tree = btoa(decodeURIComponent(encodeURIComponent(message.result.svgtree)))
        $("#svgtree").html("<img src='data:image/svg+xml;base64," + b64_tree + "'></img>");
    }
    // console.log(message);
}

function createErrorTable(errors){
    if (errors == null) {
        return;
    }
    let tbody = document.querySelector('#table-errors tbody');
    tbody.innerHTML = '';
    errors.forEach((error, index) => {
        let tr = document.createElement('tr');
        let td = document.createElement('td');
        td.innerHTML = index + 1;
        tr.appendChild(td);
        td = document.createElement('td');
        td.innerHTML = error.Description;
        tr.appendChild(td);
        td = document.createElement('td');
        td.innerHTML = error.Scope;
        tr.appendChild(td);
        td = document.createElement('td');
        td.innerHTML = error.Line;
        tr.appendChild(td);
        td = document.createElement('td');
        td.innerHTML = error.Column;
        tr.appendChild(td);
        tbody.appendChild(tr);
    });
}

function getVarType(type){
    switch (type) {
        case 0:
            return 'Int';
        case 1:
            return 'Float';
        case 2:
            return 'String';
        case 3:
            return 'Bool';
        case 4:
            return 'nil';
        case 5:
            return 'Character';
    }
}

function createSymbolTable(symbols){
    if (symbols == null) {
        return;
    }
    let tbody = document.querySelector('#table-symbols tbody');
    tbody.innerHTML = '';
    symbols.forEach(symbol => {
        let tr = document.createElement('tr');
        let td = document.createElement('td');
        td.innerHTML = symbol.Id;
        tr.appendChild(td);
        td = document.createElement('td');
        td.innerHTML = symbol.SymbolType;
        tr.appendChild(td);
        td = document.createElement('td');
        td.innerHTML = getVarType(symbol.DataType);
        tr.appendChild(td);
        td = document.createElement('td');
        td.innerHTML = symbol.Scope;
        tr.appendChild(td);
        td = document.createElement('td');
        td.innerHTML = symbol.Line;
        tr.appendChild(td);
        td = document.createElement('td');
        td.innerHTML = symbol.Col;
        tr.appendChild(td);
        tbody.appendChild(tr);
    });
}

function leerArchivo() {
    const fileInput = document.getElementById("fileInput");
    const contenidoArchivo = document.getElementById("contenidoArchivo");
  
    const archivo = fileInput.files[0];
  
    if (archivo) {
      const lector = new FileReader();
  
      lector.onload = function(evento) {
        const contenido = evento.target.result;
        contenidoArchivo.textContent = contenido;
      };
  
      lector.readAsText(archivo);
    } else {
      contenidoArchivo.textContent = "No se ha seleccionado ningún archivo.";
    }
  }

function guardarArchivo(name) {
    // Obtener el contenido del textarea
    const contenido = editor.value;
  
    // Crear un Blob con el contenido y especificar el tipo de archivo
    const archivoBlob = new Blob([contenido], { type: 'text/plain' });
  
    // Crear un objeto URL para el Blob
    const urlArchivo = window.URL.createObjectURL(archivoBlob);
  
    // Crear un enlace <a> para la descarga
    const enlaceDescarga = document.createElement('a');
    enlaceDescarga.href = urlArchivo;
    enlaceDescarga.download = name + '.swift'; // Nombre del archivo
  
    // Simular un clic en el enlace para iniciar la descarga
    enlaceDescarga.click();
  
    // Liberar el objeto URL cuando ya no sea necesario
    window.URL.revokeObjectURL(urlArchivo);
}

/* Guardar como */
save_as.addEventListener('click', (e) => {
    guardarArchivo("archivo");
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
    // console.log(message.error[0]);
    /* for (let i = 0; i < message.error.length; i++) {
        console.log(message.error[i].Scope);
    } */
    createErrorTable(message.error)
    createSymbolTable(message.symbol)
    terminal.value +=  message.message;
});

/* Limpiar termianl */
btn_clear.addEventListener('click', (e) => {
    e.preventDefault();
    terminal.value = '';
})
