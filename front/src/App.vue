<template>
    <TheLayout id="app">
        <template #branding>
            <VHead>{{projectName}}</VHead>
        </template>
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
    import {mapGetters, mapActions} from 'vuex';
    import {ProjectInfoFilter} from "../../webapp/jstypes/apiModel";
    import VButton from "swtui/src/components/VButton";
    import VHead from "swtui/src/components/VHead";

    export default {
        components: {VHead, VButton, MainMenu, TheLayout},
        methods: {
            ...mapGetters([
                'getListProjectInfo',
            ]),
            ...mapActions([
                'findProjectInfo',
            ])
        },
        data() {
            return {
                piFileter: new ProjectInfoFilter(),
            };
        },
        created() {

            this.piFileter.CurrentPage = 1;
            this.piFileter.PerPage = 1000;

            this.findProjectInfo({filter: this.piFileter});
        },
        computed: {

            projectName() {

                let infoList = this.getListProjectInfo();

                let name = infoList.find(item => item.Name === "ProjectName");

                return name ? name.Value : '';
            },
        },

    }
</script>