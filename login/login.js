// Base de usuários e senhas (simples e local)
const users = [
  { username: "admin", password: "123456" },
  { username: "user", password: "senha123" }
];

document.getElementById("loginForm").addEventListener("submit", function (e) {
  e.preventDefault();

  const inputUser = document.getElementById("username").value;
  const inputPass = document.getElementById("password").value;

  const validUser = users.find(user => user.username === inputUser && user.password === inputPass);

  if (validUser) {
    alert("Login bem-sucedido! Redirecionando...");
    window.location.href = "/quiz/quiz.html"; // redireciona para a pasta do quiz
  } else {
    alert("Usuário ou senha incorretos.");
  }
});