<template>
    <VSet>
        <WorkSpace width="fit">
            <VSet vertical>
                <VText @click="text+=text">{{text}}</VText>
                <VImage :src="src" data-bubu="task" draggable="true" v-on:dragend="drop($event)"></VImage>
            </VSet>
        </WorkSpace>
        <WorkSpace noIndent>
            <canvas id="SomeCanvas" v-on:drop="drop($event)">></canvas>
        </WorkSpace>
    </VSet>
</template>

<script>

    import BuBu from '../bubu/main';
    import {TYPES} from '../bubu/constants';
    import VLink from "swtui/src/components/VLink";
    import VButton from "swtui/src/components/VButton";
    import VSet from "swtui/src/components/VSet";
    import WorkSpace from "swtui/src/components/WorkSpace";
    import VText from "swtui/src/components/VText";
    import VImage from "swtui/src/components/VImage";

    export default {
        name: "Bpm",
        components: {VImage, VText, WorkSpace, VSet, VButton, VLink},
        data() {
            return {
                text: "Some asd",
                bubu: {},
                src: "",
            };
        },
        updated() {
            this.bubu.UpdateCanvas();
        },
        created() {


            this.$nextTick(() => {

                this.bubu = new BuBu('SomeCanvas');
                this.src = this.bubu.GetSrcImageTask();
                this.bubu.UpdateCanvas();
                this.bubu.Render();

            });

        },
        methods: {
            drop: function(event) {
                console.log('event',event);
            },
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