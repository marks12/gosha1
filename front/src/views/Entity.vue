<template>
    <EntityGen>
        <template #data>
            <table>
                <thead>
                    <tr>
                        <th>ID</th>
                        <th>Name</th>
                        <th>Fields</th>
                    </tr>
                </thead>

                <tbody>
                    <tr
                            v-for="entityItem in entityList"
                            :key="entityItem.Id"
                            @click="selectEntityItem(entityItem)"
                            class="sw-table__row_can-select"
                            :class="{'sw-table__row_is-selected': entityItem.Id === currentEntityItem.item.Id}"
                    >
                        <td>
                            {{ entityItem.Id }}
                        </td>
                        <td>
                            <VHead level="h2">{{ entityItem.Name }}</VHead>
                        </td>
                        <td>
                            <VSet v-if="entityItem.Fields" vertical>
                                <VText v-for="f in entityItem.Fields" :key="f.Id">
                                    <VSet>
                                        <VText width="col2">{{ f.Name }}</VText>
                                        <VBadge>{{ f.Type }}</VBadge>
                                    </VSet>
                                </VText>
                            </VSet>
                        </td>
                    </tr>
                </tbody>
            </table>
        </template>
    </EntityGen>
</template>

<script>
    import EntityGen from "../../../webapp/jstypes/components/EntityGen";
    import VBadge from "swtui/src/components/VBadge";
    export default {
        name: "Entity",
        components: {VBadge, EntityGen},
        mixins: [
            EntityGen,
        ],
    }
</script>

<style scoped>
    .h100 {
        height: 100vh;
    }
</style>