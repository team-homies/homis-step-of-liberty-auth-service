📦homis-boilerplate<br>
 ┣ 📂app<br>
 ┃ ┣ 📂api<br>
 ┃ ┃ ┣ 📂patient<br>
 ┃ ┃ ┃ ┣ 📂handler<br>
 ┃ ┃ ┃ ┃ ┗ 📜handler.go<br>
 ┃ ┃ ┃ ┣ 📂resource<br>
 ┃ ┃ ┃ ┃ ┣ 📜patient.go<br>
 ┃ ┃ ┃ ┃ ┣ 📜request.go<br>
 ┃ ┃ ┃ ┃ ┗ 📜response.go<br>
 ┃ ┃ ┃ ┣ 📂service<br>
 ┃ ┃ ┃ ┃ ┗ 📜service.go<br>
 ┃ ┃ ┃ ┗ 📜patient.go<br>
 ┃ ┃ ┗ 📜router.go<br>
 ┃ ┣ 📂grpc<br>
 ┃ ┃ ┣ 📂proto<br>
 ┃ ┃ ┃ ┣ 📂dosage<br>
 ┃ ┃ ┃ ┃ ┣ 📜dosage.pb.go<br>
 ┃ ┃ ┃ ┃ ┣ 📜dosage.proto<br>
 ┃ ┃ ┃ ┃ ┗ 📜dosage_grpc.pb.go<br>
 ┃ ┃ ┃ ┗ 📂patient<br>
 ┃ ┃ ┃ ┃ ┣ 📜patient.pb.go<br>
 ┃ ┃ ┃ ┃ ┣ 📜patient.proto<br>
 ┃ ┃ ┃ ┃ ┗ 📜patient_grpc.pb.go<br>
 ┃ ┃ ┣ 📂service<br>
 ┃ ┃ ┃ ┣ 📂dosage<br>
 ┃ ┃ ┃ ┃ ┗ 📜dosage.go<br>
 ┃ ┃ ┃ ┗ 📂patient<br>
 ┃ ┃ ┃ ┃ ┗ 📜patient.go<br>
 ┃ ┃ ┗ 📜grpc.go<br>
 ┃ ┗ 📜app.go<br>
 ┣ 📂common<br>
 ┃ ┣ 📂fiberkit<br>
 ┃ ┃ ┗ 📜fiberkit.go<br>
 ┃ ┗ 📂util<br>
 ┃ ┃ ┗ 📜util.go<br>
 ┣ 📂config<br>
 ┃ ┣ 📜cofig.go<br>
 ┃ ┣ 📜config.yml<br>
 ┃ ┗ 📜constant.go<br>
 ┣ 📂constant<br>
 ┃ ┣ 📂common<br>
 ┃ ┃ ┗ 📜common.go<br>
 ┃ ┣ 📂path<br>
 ┃ ┃ ┣ 📂core<br>
 ┃ ┃ ┃ ┗ 📜core.go<br>
 ┃ ┃ ┗ 📂external<br>
 ┃ ┗ 📜constant.go<br>
 ┣ 📂database<br>
 ┃ ┣ 📂entity<br>
 ┃ ┃ ┗ 📜patient.go<br>
 ┃ ┣ 📂repository<br>
 ┃ ┃ ┣ 📂patient<br>
 ┃ ┃ ┃ ┗ 📜patient.go<br>
 ┃ ┃ ┗ 📜repository.go<br>
 ┃ ┗ 📜connection.go<br>
 ┣ 📂external<br>
 ┣ 📂middleware<br>
 ┣ 📜.gitignore<br>
 ┣ 📜go.mod<br>
 ┣ 📜go.sum<br>
 ┣ 📜main.go<br>
 ┗ 📜tree<br>
