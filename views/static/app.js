const login = document.querySelector('.login');
const modalBg = document.querySelector('.modal-bg');
const modalClose = document.querySelector('.modal-close');

login.addEventListener('click', function () {
    modalBg.classList.add('bg-active');
});

modalClose.addEventListener('click', function () {
    modalBg.classList.remove('bg-active');
});

const cadastrar = document.querySelector('.cadastrar');
const modalCadastro = document.querySelector('.cadastro');
const modalCadastroClose = document.querySelector('.cadastro-close');

cadastrar.addEventListener('click', function () {
    modalCadastro.classList.add('bg-active-cadastro');
});

modalCadastroClose.addEventListener('click', function () {
    modalCadastro.classList.remove('bg-active-cadastro');
});