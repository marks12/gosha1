<template>
    <TheLayout id="app">
        <template #navigation>
            <MainMenu></MainMenu>
        </template>
        <template #content>
            <router-view #content/>
        </template>
    </TheLayout>
</template>

<style lang="scss">
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}
#nav {
  padding: 30px;
  a {
    font-weight: bold;
    color: #2c3e50;
    &.router-link-exact-active {
      color: #42b983;
    }
  }
}
</style>
<script>
    import TheLayout from "swtui/src/components/TheLayout";
    import MainMenu from "./components/MainMenu";
    export default {
        components: {MainMenu, TheLayout}
    }
</script>