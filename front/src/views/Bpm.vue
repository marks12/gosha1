<template>
    <div style="width: 100%">
        <canvas id="SomeCanvas" width="500" height="600"></canvas>
    </div>
</template>

<script>

    import BuBu from '../bubu/main';
    import {TYPES} from '../bubu/constants';
    import VLink from "swtui/src/components/VLink";
    import VButton from "swtui/src/components/VButton";
    import VSet from "swtui/src/components/VSet";

    export default {
        name: "Bpm",
        components: {VSet, VButton, VLink},
        data() {
            return {
                text: "Some asd",
                bubu: {},
            };
        },
        created() {

            setTimeout(()=>{

                this.bubu = new BuBu('SomeCanvas');
                this.bubu.Render();

            });

        },
        methods: {
            addAction() {

                this.bubu.Add(new this.bubu.Elements.Task({
                    name: "Task1 name",
                    description: "Task1 description",
                })).Render();

            },
            addCondition() {

                this.bubu.Add(
                    new this.bubu.Elements.Condition({
                        name: "Condition name",
                        description: "Condition description",
                        Coords: {
                            X: 110,
                            Y: 0,
                        },
                    })
                ).Render();
            },
        },
    }
</script>

<style scoped>

</style>