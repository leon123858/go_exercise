// Import the functions you need from the SDKs you need
import { initializeApp } from 'firebase/app';
import {
	createUserWithEmailAndPassword,
	signInWithEmailAndPassword,
	deleteUser,
	getAuth,
} from 'firebase/auth';

// TODO: Add SDKs for Firebase products that you want to use
// https://firebase.google.com/docs/web/setup#available-libraries

// Your web app's Firebase configuration
// For Firebase JS SDK v7.20.0 and later, measurementId is optional
const firebaseConfig = {
	apiKey: 'AIzaSyBML52CAtCHsKenvpbvfF3wKWiAoqJBKaM',
	authDomain: 'communityaimeeting.firebaseapp.com',
	projectId: 'communityaimeeting',
	storageBucket: 'communityaimeeting.appspot.com',
	messagingSenderId: '1034436386392',
	appId: '1:1034436386392:web:37d600a65ffa2dd251742d',
	measurementId: 'G-FS32YWTVN1',
};

// Initialize Firebase
const app = initializeApp(firebaseConfig);
const auth = getAuth(app);

// createUserWithEmailAndPassword(auth, 'test@mail.com', '000000')
// .then((_) => signInWithEmailAndPassword(auth, 'test@mail.com', '000000'))
signInWithEmailAndPassword(auth, 'test@mail.com', '000000')
	.then((pkg) => {
		return pkg.user.getIdToken();
	})
	.then((token) => {
		console.log(token);
		return token;
	});
