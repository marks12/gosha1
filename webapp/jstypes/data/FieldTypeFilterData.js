
	import {FieldTypeFilter, FieldTypeFilterFilter} from '../apiModel';

let fieldTypeFilterData = function () {
    return {
        fieldTypeFilterFilter: new FieldTypeFilterFilter(),
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
        currentFieldTypeFilterItem: {
            item: new FieldTypeFilter(),
            hasChange: false,
            isSelected: false,
            canDelete: false,
            showDeleteConfirmation: false,
        },
    }
}();

export default fieldTypeFilterData;

