
	import {APIStatus, APIStatusFilter} from '../apiModel';

let aPIStatusData = function () {
    return {
        aPIStatusFilter: new APIStatusFilter(),
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
        currentAPIStatusItem: {
            item: new APIStatus(),
            hasChange: false,
            isSelected: false,
            canDelete: false,
            showDeleteConfirmation: false,
        },
    }
}();

export default aPIStatusData;

