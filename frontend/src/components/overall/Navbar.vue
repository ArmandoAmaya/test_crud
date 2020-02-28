<template>
    <header id="header_page" class="header-page">
        <b-navbar variant="faded" type="light" class="border-bottom">
            <b-container>
                <b-navbar-brand href="#">CRUD</b-navbar-brand>
                <b-navbar-toggle target="nav-collapse"></b-navbar-toggle>
                <b-collapse id="nav-collapse" is-nav>
                    <b-navbar-nav class="ml-auto" v-if="isLoggedIn">
                        <b-nav-item href="#" @click.prevent="logOut">Cerrar sesión</b-nav-item>
                    </b-navbar-nav>
                    <b-navbar-nav class="ml-auto" v-else>
                        <b-nav-item to="/login">Iniciar sesión</b-nav-item>
                        <b-nav-item to="/register">Registro</b-nav-item>
                    </b-navbar-nav>

                </b-collapse>
            </b-container>
        </b-navbar>
    </header>
</template>

<script>
//import Logout from './Logout'
import {mapGetters} from 'vuex'

export default {
    name: 'Navbar',
    data(){
        return {}
    },
    computed: Object.assign(mapGetters([
        'isLoggedIn',
        'user'
    ]), {
            
    }),
    components: {

    },
    methods: {
        logOut(){
            this.$store.dispatch("logout");
        }
    },

    mounted(){}
    
};
</script>

<style lang="scss" scoped></style>