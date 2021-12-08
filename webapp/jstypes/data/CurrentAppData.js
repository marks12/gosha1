
	import {CurrentApp, CurrentAppFilter} from '../apiModel';

let currentAppData = function () {
    return {
        currentAppFilter: new CurrentAppFilter(),
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
        currentCurrentAppItem: {
            item: new CurrentApp(),
            hasChange: false,
            isSelected: false,
            canDelete: false,
            showDeleteConfirmation: false,
        },
    }
}();

export default currentAppData;

