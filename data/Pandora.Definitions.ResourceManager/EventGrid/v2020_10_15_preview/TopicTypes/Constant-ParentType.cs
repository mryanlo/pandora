using Pandora.Definitions.Attributes;
using System.ComponentModel;

namespace Pandora.Definitions.ResourceManager.EventGrid.v2020_10_15_preview.TopicTypes
{
    [ConstantType(ConstantTypeAttribute.ConstantType.String)]
    internal enum ParentTypeConstant
    {
        [Description("domains")]
        Domains,

        [Description("topics")]
        Topics,
    }
}
