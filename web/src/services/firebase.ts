import { initializeApp } from "firebase/app";
import {
  browserLocalPersistence,
  getAuth,
  setPersistence
} from "firebase/auth";
import { type App } from "vue";

// Firebase configuration
const firebaseConfig = {
  apiKey: "AIzaSyC8zV25C82rxrZN_gQOOGV_OmTX2akQP4c",
  authDomain: "badminton-stats-3e3e.firebaseapp.com",
  projectId: "badminton-stats-3e3e",
  storageBucket: "badminton-stats-3e3e.firebasestorage.app",
  messagingSenderId: "754749631741",
  appId: "1:754749631741:web:53b7ca441da7099deae294",
};

// Initialize Firebase App
const firebaseApp = initializeApp(firebaseConfig);
const auth = getAuth(firebaseApp);

// Plugin Installation
export default {
  install(app: App) {
    // Set up Firebase Auth persistence
    setPersistence(auth, browserLocalPersistence)
      .then(() => {
        console.log("Firebase persistence set to local.");
      })
      .catch((error) => {
        console.error("Failed to set Firebase persistence:", error);
      });

    // Provide Firebase Auth instance globally
    app.provide("auth", auth);
  },
};

// Export Firebase instances for direct use if needed
export { auth, firebaseApp };
