<template>
  <div>

    <div class="sidenav">
       <div class="login-main-text">
          <h2>Application<br> Login Page</h2>
          <p>Login or register from here to access.</p>
       </div>
    </div>

    <div class="main">
       <div class="col-md-6 col-sm-12" >

          <div class="login-form">

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

              <form @submit.prevent="login" id="loginform">
                <b-form-group label="Email:" label-for="id_email">
                  <b-form-input
                    id="id_email"
                    required 
                    placeholder="Ej: email@demo.com" 
                    type="email" 
                    name="email" 
                  ></b-form-input>
                </b-form-group>


                <b-form-group label="Contraseña:" label-for="id_pass">
                  <b-form-input
                    id="id_pass"
                    required 
                    placeholder="Contraseña de ingreso" 
                    type="password" 
                    name="password" 
                  ></b-form-input>
                </b-form-group>
               
                <!-- <button type="submit" class="btn btn-black text-white mr-1"> Ingresar</button> -->
                <vue-ladda class="btn btn-black text-white mr-1" :loading="loading">Ingresar</vue-ladda>
                <router-link to="/register" class="btn btn-secondary">Registrarse</router-link>
             </form>
          </div>
       </div>
    </div>

    <!-- <form class="login" @submit.prevent="login" id="loginform">
      <h1>Sign in</h1>
      <label>Email</label>
      <input required name="email" type="email" placeholder="Name" />
      <label>Password</label>
      <input
        required
        name="password"
        type="password"
        placeholder="Password"
      />
      <hr />
      <button type="submit">Login</button>
    </form> -->
  </div>
</template>

<script>
export default {
  name: "Login",
  data() {
    return {
      loading: false,
      dismissSecs: 10,
      dismissCountDown: 0,
      showDismissibleAlert: false,
      type: 'warning',
      alertMessage: ''
    };
  },
  methods: {
    login() {
      let myForm = document.getElementById("loginform");
      let formData = new FormData(myForm);
      this.loading = true;

      this.$store
        .dispatch("login", formData)
        .then((resp) => {
          this.showAlert('success');
          this.alertMessage = resp.data.message;
          var $this = this;
          setTimeout(function(){
            $this.$router.push("/");
          }, 2000);
          this.loading = false;
        })
        .catch(err => {
          this.showAlert('danger');
          this.alertMessage = err;
          this.loading = false;
          //console.log(err)
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
  mounted() {
      document.getElementById('header_page').classList.add('d-none');
  },
  beforeDestroy() {
      document.getElementById('header_page').classList.remove('d-none');
  }
};
</script>

<style lang="scss" scoped></style>
