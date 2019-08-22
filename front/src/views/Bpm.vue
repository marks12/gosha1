<template>
    <div>
        Canvas become
        {{text}}
    </div>
</template>

<script>

    import BuBu from '../bubu/main';

    export default {
        name: "Bpm",
        data() {
            return {
                text: "Some asd",
                bubu: new BuBu(),
            };
        },
        created() {

            this.bubu.Add(
                this.bubu.Elements.CrAction()
                    .SetName("Action1 name")
                    .SetDescription("action 1 description")
            );

            this.bubu.Add(
                this.bubu.Elements.CrCondition()
                    .SetName("Condition1")
            );

            console.log('names', this.bubu.GetElementsByName("Condition1"));

        },
    }
</script>

<style scoped>

</style>