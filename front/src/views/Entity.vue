<template>
    <WorkSpace width="dyn">
        <template #header>
            <VHead level="h1" align="left">Entity</VHead>
        </template>
        <template #content>
            <table slot="content" cellpadding="5" cellspacing="5">
                <thead>
                    <tr>
                        <th>Name</th>
                        <th>Fields</th>
                    </tr>
                </thead>
                <tbody>
                    <tr>
                        <td>Name</td>
                        <td>Fields</td>
                    </tr>
                </tbody>
            </table>
        </template>
        <template #footer>
            <VSet>
                <VButton text="Создать"></VButton>
            </VSet>
        </template>
    </WorkSpace>
</template>

<script>
    import WorkSpace from "swtui/src/components/WorkSpace";
    import VHead from "swtui/src/components/VHead";
    import VButton from "swtui/src/components/VButton";
    import VSet from "swtui/src/components/VSet";
    export default {
        name: "Entity",
        components: {VSet, VButton, VHead, WorkSpace}
    }
</script>

<style scoped>
    .h100 {
        height: 100vh;
    }
</style>