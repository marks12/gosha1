
	import {EntityFieldFilter, EntityFieldFilterFilter} from '../apiModel';

let entityFieldFilterData = function () {
    return {
        entityFieldFilterFilter: new EntityFieldFilterFilter(),
        panel: {
            type: '',
            show: false,
            create: 'create',
            edit: 'edit'
        },
        panelHeaderCreate: 'Создать',
        panelHeaderEdit: 'Изменить',
        panelSubmitButtonTextCreate: 'Создать',
        panelSubmitButtonTextEdit: 'Изменить',
        currentEntityFieldFilterItem: {
            item: new EntityFieldFilter(),
            hasChange: false,
            isSelected: false,
            canDelete: false,
            showDeleteConfirmation: false,
        },
    }
}();

export default entityFieldFilterData;

