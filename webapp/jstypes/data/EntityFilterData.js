
	import {EntityFilter, EntityFilterFilter} from '../apiModel';

let entityFilterData = function () {
    return {
        entityFilterFilter: new EntityFilterFilter(),
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
        currentEntityFilterItem: {
            item: new EntityFilter(),
            hasChange: false,
            isSelected: false,
            canDelete: false,
            showDeleteConfirmation: false,
        },
    }
}();

export default entityFilterData;

