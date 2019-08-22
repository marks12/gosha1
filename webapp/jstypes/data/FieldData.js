
	import {Field, FieldFilter} from '../apiModel';

let fieldData = function () {
    return {
        fieldFilter: new FieldFilter(),
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
        currentFieldItem: {
            item: new Field(),
            hasChange: false,
            isSelected: false,
            canDelete: false,
            showDeleteConfirmation: false,
        },
    }
}();

export default fieldData;

