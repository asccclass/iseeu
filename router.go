// router.go
package main

import(
   "os"
   "github.com/gorilla/mux"
   "github.com/asccclass/serverstatus"
   "github.com/asccclass/staticfileserver"
   // "github.com/asccclass/staticfileserver/libs/socketio"
   "github.com/asccclass/staticfileserver/libs/webrtc"
   "github.com/asccclass/staticfileserver/libs/qrcode"
)

// Create your Router function
func NewRouter(srv *SherryServer.ShryServer, documentRoot string)(*mux.Router) {
   router := mux.NewRouter()

   // webRTC
   wrtc, _ := SherryWebRTC.NewSryWebRTC()
   wrtc.AddRouter(router)

   // socketio
   srv.Socketio.AddRouter(router)

   // QRCode
   if os.Getenv("QRCodePath") != "" {
      // var QRCode  *qrcodeGeneratorService.QRCodeGenerator
      // var err error
      QRCode, err := qrcodeGeneratorService.NewQRCodeGenerator(srv, os.Getenv("QRCodePath"))
      if err == nil {
         QRCode.AddRouter(router)
      }
   }

   //logger
   router.Use(SherryServer.ZapLogger(srv.Logger))

   // health check
   systemName := os.Getenv("SystemName")
   m := serverstatus.NewServerStatus(systemName)
   router.HandleFunc("/healthz", m.Healthz).Methods("GET")

   // Static File server
   staticfileserver := SherryServer.StaticFileServer{documentRoot, "index.html"}
   router.PathPrefix("/").Handler(staticfileserver)

   return router
}
