<template>
    <div>
        <b-container class="mt-5">
            <b-form id="usuarios_form" @submit.prevent="create">
                <b-card>
                    <b-card-title>Crear nuevo usuario</b-card-title>
                    <b-alert
                        :show="dismissCountDown"
                        dismissible
                        :variant="type"
                        @dismissed="dismissCountDown=0"
                        @dismiss-count-down="countDownChanged"
                    >
                        <p>{{ alertMessage }}</p>
                        <b-progress
                            variant="warning"
                            :max="dismissSecs"
                            :value="dismissCountDown"
                            height="4px"
                        ></b-progress>
                    </b-alert>
                    <b-row>
                        <b-col md="6">
                            <b-form-group
                                label="Nombre:"
                                label-for="id_name"
                              >
                                <b-form-input
                                  id="id_name"
                                  type="text"
                                  required
                                  placeholder="Nombre del usuario"
                                  name="name"
                                  v-model="form.name"
                                ></b-form-input>
                            </b-form-group>
                        </b-col>

                        <b-col md="6">
                            <b-form-group
                                label="Correo electrónico:"
                                label-for="id_email"
                              >
                                <b-form-input
                                  id="id_email"
                                  type="email"
                                  required
                                  placeholder="Email del usuario"
                                  name="email"
                                  v-model="form.email"
                                ></b-form-input>
                            </b-form-group>
                        </b-col>

                        <b-col md="6">
                            <b-form-group
                                label="Contraseña:"
                                label-for="id_pass"
                              >
                                <b-form-input
                                  id="id_pass"
                                  type="password"
                                  placeholder="Contrase de ingreso"
                                  name="password"
                                ></b-form-input>
                            </b-form-group>
                        </b-col>

                        <b-col md="6">
                            <b-form-group
                                label="Repetir contraseña:"
                                label-for="id_pass_repeat"
                              >
                                <b-form-input
                                  id="id_pass_repeat"
                                  type="password"
                                  placeholder="Vuelve a escribir la contraseña"
                                  name="password_repeat"
                                ></b-form-input>
                            </b-form-group>
                        </b-col>

                    </b-row>
                
                    
                    <template v-slot:footer>
                        <vue-ladda class="btn btn-outline-primary" :loading="loading">Guardar</vue-ladda>
                    </template>
                    
                </b-card>
            </b-form>
        </b-container>
    </div>
</template>

<script>
import axios from "axios";
export default {
    name: "editar_usuario",
    data(){
        return {
            loading: false,
            dismissSecs: 10,
            dismissCountDown: 0,
            showDismissibleAlert: false,
            type: 'warning',
            alertMessage: '',
            form : {
                email: null,
                name: null
            },
            id_user: null
        }
    },
    methods: {
        create(){
            let myForm = document.getElementById("usuarios_form");
            let formData = new FormData(myForm);
            this.loading = true;

            axios
            .post(this.$store.state.api + "usuarios/edit/"+this.id_user, formData, {
                headers: {
                    "Content-Type": "multipart/form-data"
                }
            })
            .then(resp => {
                var data = resp.data;
                if (true == data.success) {
                    this.showAlert('success');
                    this.alertMessage = data.message;
                    var $this = this;
                    setTimeout(function(){
                        $this.$router.push("/");
                    }, 2000);
                    this.loading = false;
                }else{
                    this.showAlert('danger');
                    this.alertMessage = data.message;
                    this.loading = false;
                }
                
            });
        },
        countDownChanged(dismissCountDown) {
            this.dismissCountDown = dismissCountDown
        },
        showAlert(typeAlert) {
            this.type = typeAlert;
            this.dismissCountDown = this.dismissSecs
        }
    },
    mounted(){
        this.id_user = this.$route.params.id;

        axios
        .get(this.$store.state.api + "usuarios/"+this.id_user, {
            headers: {
                "Content-Type": "multipart/form-data"
            }
        })
        .then(resp => {
            var data = resp.data;
            if (false == data.success) {
               this.$router.push('/');
            }else{
                this.form.email = data.email;
                this.form.name = data.name;
            }
        });
    }
};
</script>

<style lang="scss" scoped></style>
