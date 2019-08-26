<template>
    <div style="width: 100%">
        <canvas id="SomeCanvas" width="500" height="600"></canvas>
    </div>
</template>

<script>

    import BuBu from '../bubu/main';
    import {TYPES} from '../bubu/constants';

    export default {
        name: "Bpm",
        data() {
            return {
                text: "Some asd",
                bubu: {},
            };
        },
        created() {

            setTimeout(()=>{

                this.bubu = new BuBu('SomeCanvas');

                let action = new this.bubu.Elements.Action({
                    name: "Action1 name",
                    description: "Action1 description",
                });
                this.bubu.Add(action);

                let Condition = new this.bubu.Elements.Condition({
                    name: "Condition name",
                    description: "Condition description",
                    Coords: {
                        X: 110,
                        Y: 0,
                    },
                });
                this.bubu.Add(Condition);

                console.log('names', this.bubu.GetElementsByName("Condition name"));
                console.log('types', this.bubu.GetElementsByType(TYPES.CONDITION));
                console.log('names', this.bubu.GetNames());
                console.log('found element', action);

                this.bubu.Render();

            });

        },
    }
</script>

<style scoped>

</style>