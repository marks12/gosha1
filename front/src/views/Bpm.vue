<template>
    <div>
        Canvas become
        {{text}}
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
                bubu: new BuBu(),
            };
        },
        created() {

            let action = new this.bubu.Elements.Action({
                name: "Action1 name",
                description: "Action1 description",
                Coords: {
                    X: 2,
                    Y: 3,
                },
            });
            this.bubu.Add(action);

            let Condition = new this.bubu.Elements.Condition({
                name: "Condition name",
                description: "Condition description",
            });
            this.bubu.Add(Condition);

            console.log('names', this.bubu.GetElementsByName("Condition name"));
            console.log('types', this.bubu.GetElementsByType(TYPES.CONDITION));
            console.log('names', this.bubu.GetNames());
            console.log('found element', action.Coords.GetX());
            console.log('found element', action.Coords.GetY());
        },
    }
</script>

<style scoped>

</style>