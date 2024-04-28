importScripts("https://www.gstatic.com/firebasejs/8.9.1/firebase-app.js");
importScripts("https://www.gstatic.com/firebasejs/8.9.1/firebase-messaging.js");

firebase.initializeApp({
    apiKey: "AIzaSyAaiO2DaatutvhA2b3db7CmKVnns6DXN5E",
    authDomain: "testLogin-a6302.firebaseapp.com",
    projectId: "testLogin-a6302",
    storageBucket: "testLogin-a6302.appspot.com",
    messagingSenderId: "714089444314",
    appId: "1:714089444314:web:91bd215a3d00165b1806e3",
    measurementId: "G-6C23Y0FYCR"
});
const messaging = firebase.messaging();

// const config = {
//     apiKey: 'AIzaSyDK3HZ2EJTI_3N2ya26MFQ8xLfj0v6ab78',
//     authDomain: 'test2023-45324.firebaseapp.com',
//     databaseURL: 'https://test2023-45324.firebaseio.com',
//     projectId: 'test2023-45324',
//     storageBucket: 'test2023-45324.appspot.com',
//     messagingSenderId: '153149032163'
// };
// firebase.initializeApp(config)

// messaging.onBackgroundMessage((payload) => {
//   console.log('[firebase-messaging-sw.js] Received background message ', payload);
//   // Customize notification here
//   const notificationTitle = payload.notification.title;
//   const notificationOptions = {
//     body: payload.notification.body
//   };

//   self.registration.showNotification(notificationTitle,
//     notificationOptions);
// });