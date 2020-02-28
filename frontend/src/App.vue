<template>
  <div>
    <Topnav></Topnav>
    <router-view></router-view>
  </div>
  
</template>

<script>
import { mapGetters } from "vuex";
import Topnav from './components/overall/Navbar.vue'

export default {
  name: "App",
  computed: Object.assign(mapGetters(['isLoggedIn']), {}),
  components: { Topnav },
  mounted() {
    if(this.isLoggedIn) {
      this.$store.dispatch('getOwnerUser').catch(() => {
        this.$store.dispatch('logout');
        this.$router.push('/login');
      });
    }
  }
};
</script>

<style lang="scss"></style>
