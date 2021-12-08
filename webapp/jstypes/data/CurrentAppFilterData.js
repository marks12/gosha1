
	import {CurrentAppFilter, CurrentAppFilterFilter} from '../apiModel';

let currentAppFilterData = function () {
    return {
        currentAppFilterFilter: new CurrentAppFilterFilter(),
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
        currentCurrentAppFilterItem: {
            item: new CurrentAppFilter(),
            hasChange: false,
            isSelected: false,
            canDelete: false,
            showDeleteConfirmation: false,
        },
    }
}();

export default currentAppFilterData;

