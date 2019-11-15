
	import {ResourceFilter, ResourceFilterFilter} from '../apiModel';

let resourceFilterData = function () {
    return {
        resourceFilterFilter: new ResourceFilterFilter(),
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
        currentResourceFilterItem: {
            item: new ResourceFilter(),
            hasChange: false,
            isSelected: false,
            canDelete: false,
            showDeleteConfirmation: false,
        },
    }
}();

export default resourceFilterData;

