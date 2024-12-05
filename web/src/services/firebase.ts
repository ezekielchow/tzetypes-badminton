// Import the functions you need from the SDKs you need
import { initializeApp } from "firebase/app";
import { getAuth } from "firebase/auth";
// TODO: Add SDKs for Firebase products that you want to use
// https://firebase.google.com/docs/web/setup#available-libraries

// Your web app's Firebase configuration
const firebaseConfig = {
  apiKey: "AIzaSyC8zV25C82rxrZN_gQOOGV_OmTX2akQP4c",
  authDomain: "badminton-stats-3e3e.firebaseapp.com",
  projectId: "badminton-stats-3e3e",
  storageBucket: "badminton-stats-3e3e.firebasestorage.app",
  messagingSenderId: "754749631741",
  appId: "1:754749631741:web:53b7ca441da7099deae294"
};

// Initialize Firebase
const app = initializeApp(firebaseConfig);
export const auth = getAuth(app);
