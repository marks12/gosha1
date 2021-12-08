
	import {APIError, APIErrorFilter} from '../apiModel';

let aPIErrorData = function () {
    return {
        aPIErrorFilter: new APIErrorFilter(),
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
        currentAPIErrorItem: {
            item: new APIError(),
            hasChange: false,
            isSelected: false,
            canDelete: false,
            showDeleteConfirmation: false,
        },
    }
}();

export default aPIErrorData;

