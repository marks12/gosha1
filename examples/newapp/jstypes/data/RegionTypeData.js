
	import {RegionType, RegionTypeFilter} from '../apiModel';

let regionTypeData = function () {
    return {
        regionTypeFilter: new RegionTypeFilter(),
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
        currentRegionTypeItem: {
            item: new RegionType(),
            hasChange: false,
            isSelected: false,
            canDelete: false,
            showDeleteConfirmation: false,
        },
    }
}();

export default regionTypeData;

