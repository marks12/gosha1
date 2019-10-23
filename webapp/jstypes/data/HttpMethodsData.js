
	import {HttpMethods, HttpMethodsFilter} from '../apiModel';

let httpMethodsData = function () {
    return {
        httpMethodsFilter: new HttpMethodsFilter(),
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
        currentHttpMethodsItem: {
            item: new HttpMethods(),
            hasChange: false,
            isSelected: false,
            canDelete: false,
            showDeleteConfirmation: false,
        },
    }
}();

export default httpMethodsData;

