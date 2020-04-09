
	import {FieldType, FieldTypeFilter} from '../apiModel';

let fieldTypeData = function () {
    return {
        fieldTypeFilter: new FieldTypeFilter(),
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
        currentFieldTypeItem: {
            item: new FieldType(),
            hasChange: false,
            isSelected: false,
            canDelete: false,
            showDeleteConfirmation: false,
        },
    }
}();

export default fieldTypeData;

