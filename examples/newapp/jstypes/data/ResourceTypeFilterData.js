
	import {ResourceTypeFilter, ResourceTypeFilterFilter} from '../apiModel';

let resourceTypeFilterData = function () {
    return {
        resourceTypeFilterFilter: new ResourceTypeFilterFilter(),
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
        currentResourceTypeFilterItem: {
            item: new ResourceTypeFilter(),
            hasChange: false,
            isSelected: false,
            canDelete: false,
            showDeleteConfirmation: false,
        },
    }
}();

export default resourceTypeFilterData;

