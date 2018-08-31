<template>
  <v-app>
    <v-toolbar fixed app>
      <router-link :to="{name: 'home'}">Home</router-link>
      <v-toolbar-title/>
    </v-toolbar>
    <v-content>
      <router-view/>
    </v-content>
  </v-app>
</template>

<script>

export default {
  name: 'App',
  data () {
    return {
      clipped: false,
      drawer: true,
      fixed: false,
      items: [{
        icon: 'bubble_chart',
        title: 'Inspire'
      }],
      miniVariant: false,
      right: true,
      rightDrawer: false,
      title: 'Goflippy management console'
    }
  }
}
</script>
