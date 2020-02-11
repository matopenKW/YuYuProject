// Your web app's Firebase configuration
var firebaseConfig = {
	apiKey: "AIzaSyA4ryg24no1dTQks7QQwtl0G6KLd1LbOJs",
	authDomain: "mac-001-1e4f9.firebaseapp.com",
	databaseURL: "https://mac-001-1e4f9.firebaseio.com",
	projectId: "mac-001-1e4f9",
	storageBucket: "mac-001-1e4f9.appspot.com",
	messagingSenderId: "654252082570",
	appId: "1:654252082570:web:dae673a6bad01f3ed87473"
};

$(function(){

	$(document).on('click', '#btnLogin', function(){
		aush();
	});
	
	// Initialize Firebase
	firebase.initializeApp(firebaseConfig);

});

function aush(){

	var email = $('#mailAddress').val();
	var password = $('#password').val();

	firebase.auth().signInWithEmailAndPassword(email, password)
	.then(function(user) {
		login(user);
	})
	.catch(firebaseErrHandring);
}

function login(user) {
	$('#uid').val(user.uid);
	$('#loginForm').submit();
}

function firebaseErrHandring(error){
	var errorCode = error.code;
	var errorMessage = error.message;

	if (errorCode === 'auth/invalid-email') {
		alert('メールアドレスの形式が不正です。');

	} else if (errorCode === 'auth/wrong-password') {
		alert('パスワードが間違っている又は不正な形式です。');

	} else if (errorCode === 'auth/user-not-found') {
		alert('存在しないユーザー又は削除された可能性があります。');

	} else if (errorCode === 'auth/email-already-in-use'){
		alert('既に登録してあるメールアドレスです。');

	} else if (errorCode === 'auth/weak-password') {
		alert('パスワードは６桁以上で登録してください。');

	} else {
		  alert(errorMessage);

	}

	console.log(error);

}