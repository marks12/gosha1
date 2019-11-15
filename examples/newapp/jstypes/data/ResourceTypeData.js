
	import {ResourceType, ResourceTypeFilter} from '../apiModel';

let resourceTypeData = function () {
    return {
        resourceTypeFilter: new ResourceTypeFilter(),
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
        currentResourceTypeItem: {
            item: new ResourceType(),
            hasChange: false,
            isSelected: false,
            canDelete: false,
            showDeleteConfirmation: false,
        },
    }
}();

export default resourceTypeData;

