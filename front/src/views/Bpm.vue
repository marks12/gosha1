<template>
    <VSet>
        <WorkSpace width="fit">
            <VSet vertical>
                <VHead level="h2">Элементы</VHead>
                <VImage :src="src.task" data-bubu="task" draggable="true" v-on:dragend="drop($event)"></VImage>
                <VImage :src="src.condition" data-bubu="task" draggable="true" v-on:dragend="drop($event)"></VImage>
            </VSet>
        </WorkSpace>
        <WorkSpace noIndent>
            <canvas id="SomeCanvas"></canvas>
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
    import VHead from "swtui/src/components/VHead";

    export default {
        name: "Bpm",
        components: {VHead, VImage, VText, WorkSpace, VSet, VButton, VLink},
        data() {
            return {
                text: "Some asd",
                bubu: {},
                src: {
                    task: "",
                    condition: "",
                },
            };
        },
        updated() {
            this.bubu.UpdateCanvas();
        },
        created() {

            this.$nextTick(() => {

                this.bubu = new BuBu('SomeCanvas');
                this.src.task = this.bubu.GetSrcImageTask();
                this.src.condition = this.bubu.GetSrcImageCondition();
                this.bubu.UpdateCanvas();
                this.bubu.Render();
            });
        },
        methods: {
            allowDrop: function(ev) {
                ev.preventDefault();
            },
            drop: function(event) {
                this.bubu.DropElement(event)
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