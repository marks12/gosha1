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


    <!--<template>-->
        <!--<VSet vertical>-->
            <!--<VGroup width="dyn" control color="bg-light" selected>-->
                <!--<VSet  vertical indent-size="M">-->
                    <!--<VHead>Заголовок</VHead>-->
                    <!--<VSet indent-size="L" wrap>-->
                        <!--<VSet vertical indent-size="XS" width="fit">-->
                            <!--<VText width="fit">Группа</VText >-->
                            <!--<VBadge>Бэйдж</VBadge>-->
                        <!--</VSet>-->
                        <!--<VSet vertical indent-size="XS" width="fit">-->
                            <!--<VText>Гр hgf hgf h hgf hfg уппа</VText>-->
                            <!--<VBadge>Бэйдж</VBadge>-->
                        <!--</VSet>-->
                        <!--<VSet vertical indent-size="XS" width="fit">-->
                            <!--<VText>Группа</VText>-->
                            <!--<VBadge>Бэйдж</VBadge>-->
                        <!--</VSet>-->
                        <!--<VSet vertical indent-size="XS" width="fit">-->
                            <!--<VText>Группа</VText>-->
                            <!--<VBadge>Бэйдж</VBadge>-->
                        <!--</VSet>          <VSet vertical indent-size="XS" width="fit">-->
                        <!--<VText>Группа</VText>-->
                        <!--<VBadge>Бэйдж</VBadge>-->
                    <!--</VSet>-->
                        <!--<VSet vertical indent-size="XS" width="fit">-->
                            <!--<VText>Группа</VText>-->
                            <!--<VBadge>Бэйдж</VBadge>-->
                        <!--</VSet>          <VSet vertical indent-size="XS" width="fit">-->
                        <!--<VText>Группа</VText>-->
                        <!--<VBadge>Бэйдж</VBadge>-->
                    <!--</VSet>-->
                        <!--<VSet vertical indent-size="XS" width="fit">-->
                            <!--<VText>Группа</VText>-->
                            <!--<VBadge>Бэйдж</VBadge>-->
                        <!--</VSet>          <VSet vertical indent-size="XS" width="fit">-->
                        <!--<VText>Группа</VText>-->
                        <!--<VBadge>Бэйдж</VBadge>-->
                    <!--</VSet>-->
                        <!--<VSet vertical indent-size="XS" width="fit">-->
                            <!--<VText>Группа</VText>-->
                            <!--<VBadge>Бэйдж</VBadge>-->
                        <!--</VSet>          <VSet vertical indent-size="XS" width="fit">-->
                        <!--<VText>Группа</VText>-->
                        <!--<VBadge>Бэйдж</VBadge>-->
                    <!--</VSet>-->
                        <!--<VSet vertical indent-size="XS" width="fit">-->
                            <!--<VText>Группа</VText>-->
                            <!--<VBadge>Бэйдж</VBadge>-->
                        <!--</VSet>-->
                    <!--</VSet>-->
                <!--</VSet>-->
            <!--</VGroup>-->
            <!--<VGroup>-->
                <!--<VSet  vertical indent-size="S">-->
                    <!--<VHead>Заголовок</VHead>-->
                    <!--<VSet indent-size="L">-->
                        <!--<VSet vertical indent-size="XS">-->
                            <!--<VText>Группа</VText>-->
                            <!--<VBadge>Бэйдж</VBadge>-->
                        <!--</VSet>-->
                        <!--<VSet vertical indent-size="XS">-->
                            <!--<VText>Группа</VText>-->
                            <!--<VBadge>Бэйдж</VBadge>-->
                        <!--</VSet>-->
                        <!--<VSet vertical indent-size="XS">-->
                            <!--<VText>Группа</VText>-->
                            <!--<VBadge>Бэйдж</VBadge>-->
                        <!--</VSet>-->
                        <!--<VSet vertical indent-size="XS">-->
                            <!--<VText>Группа</VText>-->
                            <!--<VBadge>Бэйдж</VBadge>-->
                        <!--</VSet>-->
                    <!--</VSet>-->
                <!--</VSet>-->
            <!--</VGroup>-->

            <!--<VGroup>-->
                <!--<VSet  vertical indent-size="S">-->
                    <!--<VHead>Заголовок</VHead>-->
                    <!--<VSet indent-size="L">-->
                        <!--<VSet vertical indent-size="XS">-->
                            <!--<VText>Группа</VText>-->
                            <!--<VBadge>Бэйдж</VBadge>-->
                        <!--</VSet>-->
                        <!--<VSet vertical indent-size="XS">-->
                            <!--<VText>Группа</VText>-->
                            <!--<VBadge>Бэйдж</VBadge>-->
                        <!--</VSet>-->
                        <!--<VSet vertical indent-size="XS">-->
                            <!--<VText>Группа</VText>-->
                            <!--<VBadge>Бэйдж</VBadge>-->
                        <!--</VSet>-->
                        <!--<VSet vertical indent-size="XS">-->
                            <!--<VText>Группа</VText>-->
                            <!--<VBadge>Бэйдж</VBadge>-->
                        <!--</VSet>-->
                    <!--</VSet>-->
                <!--</VSet>-->
            <!--</VGroup>-->

        <!--</VSet>-->
    <!--</template>-->


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