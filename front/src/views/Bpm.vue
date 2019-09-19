<template>
    <VSet>
        <WorkSpace width="fit">
            <VSet vertical>
                <VHead>Элементы</VHead>
                <div id="BubuToolbox"></div>
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
                this.bubu = new BuBu('SomeCanvas', 'BubuToolbox');
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
        },
    }
</script>

<style scoped>

</style>