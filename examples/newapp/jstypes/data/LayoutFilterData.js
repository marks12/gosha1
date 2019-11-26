
	import {LayoutFilter, LayoutFilterFilter} from '../apiModel';

let layoutFilterData = function () {
    return {
        layoutFilterFilter: new LayoutFilterFilter(),
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
        currentLayoutFilterItem: {
            item: new LayoutFilter(),
            hasChange: false,
            isSelected: false,
            canDelete: false,
            showDeleteConfirmation: false,
        },
    }
}();

export default layoutFilterData;

