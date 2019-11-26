
	import {LayoutContentFilter, LayoutContentFilterFilter} from '../apiModel';

let layoutContentFilterData = function () {
    return {
        layoutContentFilterFilter: new LayoutContentFilterFilter(),
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
        currentLayoutContentFilterItem: {
            item: new LayoutContentFilter(),
            hasChange: false,
            isSelected: false,
            canDelete: false,
            showDeleteConfirmation: false,
        },
    }
}();

export default layoutContentFilterData;

