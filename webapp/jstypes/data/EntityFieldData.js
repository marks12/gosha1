
	import {EntityField, EntityFieldFilter} from '../apiModel';

let entityFieldData = function () {
    return {
        entityFieldFilter: new EntityFieldFilter(),
        panel: {
            type: '',
            show: false,
            create: 'create',
            edit: 'edit',
            request: 'request'
        },
        panelHeaderCreate: 'Создать',
        panelHeaderEdit: 'Изменить',
        panelSubmitButtonTextCreate: 'Создать',
        panelSubmitButtonTextEdit: 'Изменить',
        currentEntityFieldItem: {
            item: new EntityField(),
            hasChange: false,
            isSelected: false,
            canDelete: false,
            showDeleteConfirmation: false,
        },
    }
}();

export default entityFieldData;

