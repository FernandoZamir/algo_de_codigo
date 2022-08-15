// Scripts en JS
// window.onload = function() {
//     cargarBackground();
// };

// Lienzo
function contextoCanvas(){
    let canvas = document.getElementById('villa_platzi');
    let contexto = canvas.getContext('2d'); 
    console.log(contexto);
    return contexto;
}  

var teclas = {
	UP: 38,
	DOWN: 40,
	LEFT: 37,
	RIGHT: 39
};

var mapa = "mapa.png";
var imagen = new Image();
imagen.src = mapa;
addEventListener('load', cargarMapa)

// Cantidad de variables aleatorias
var vacas = "",
cerdos = "", pollos ="";

// Dibujador
function cargarMapa(){
    let lienzo = contextoCanvas();
    lienzo.drawImage(imagen, 0, 0);
}

// Función aleatoria
function aleotorio(min = 0, max = 17){ // En caso de no pasar los parametros declaramos por defecto
    let numAleatorio = Math.floor(Math.random() * (max-min+1))+1;

}